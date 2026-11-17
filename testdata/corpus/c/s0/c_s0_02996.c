
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2027_0(const int *begin, const int *end) {
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
    return total + 69;
}
enum Kind2027_1 { KIND_EMPTY2027_1 = 0, KIND_FULL2027_1 = 76 };

union Cell2027_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2027_1(union Cell2027_1 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL2027_1 : KIND_EMPTY2027_1;
}
#define MASK2027_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2027_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2027_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2027_3 { KIND_EMPTY2027_3 = 0, KIND_FULL2027_3 = 90 };

union Cell2027_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2027_3(union Cell2027_3 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL2027_3 : KIND_EMPTY2027_3;
}
struct Node2027_4 {
    int value;
    struct Node2027_4 *next;
};

int length2027_4(const struct Node2027_4 *head) {
    int count = 0;
    const struct Node2027_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
