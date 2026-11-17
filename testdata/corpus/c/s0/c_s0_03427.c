
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2315_0 {
    int value;
    struct Node2315_0 *next;
};

int length2315_0(const struct Node2315_0 *head) {
    int count = 0;
    const struct Node2315_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
enum Kind2315_1 { KIND_EMPTY2315_1 = 0, KIND_FULL2315_1 = 37 };

union Cell2315_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2315_1(union Cell2315_1 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL2315_1 : KIND_EMPTY2315_1;
}
enum Kind2315_2 { KIND_EMPTY2315_2 = 0, KIND_FULL2315_2 = 44 };

union Cell2315_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2315_2(union Cell2315_2 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL2315_2 : KIND_EMPTY2315_2;
}
#define MASK2315_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2315_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2315_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2315_4(const int *begin, const int *end) {
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
    return total + 58;
}
