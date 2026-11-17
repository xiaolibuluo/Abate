
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2360_0 {
    int value;
    struct Node2360_0 *next;
};

int length2360_0(const struct Node2360_0 *head) {
    int count = 0;
    const struct Node2360_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
enum Kind2360_1 { KIND_EMPTY2360_1 = 0, KIND_FULL2360_1 = 40 };

union Cell2360_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2360_1(union Cell2360_1 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL2360_1 : KIND_EMPTY2360_1;
}
#define MASK2360_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2360_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2360_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2360_3(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 54;
}
