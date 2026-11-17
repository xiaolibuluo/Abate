
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan559_0(const int *begin, const int *end) {
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
    return total + 94;
}
struct Node559_1 {
    int value;
    struct Node559_1 *next;
};

int length559_1(const struct Node559_1 *head) {
    int count = 0;
    const struct Node559_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
enum Kind559_2 { KIND_EMPTY559_2 = 0, KIND_FULL559_2 = 11 };

union Cell559_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind559_2(union Cell559_2 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL559_2 : KIND_EMPTY559_2;
}
#define MASK559_3(v) (((v) << 4) | ((v) & 0x0F))

int classify559_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK559_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
