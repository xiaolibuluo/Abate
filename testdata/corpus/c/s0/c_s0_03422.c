
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2313_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2313_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2313_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2313_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2313_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2313_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2313_2(const int *begin, const int *end) {
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
    return total + 18;
}
enum Kind2313_3 { KIND_EMPTY2313_3 = 0, KIND_FULL2313_3 = 25 };

union Cell2313_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2313_3(union Cell2313_3 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL2313_3 : KIND_EMPTY2313_3;
}
struct Node2313_4 {
    int value;
    struct Node2313_4 *next;
};

int length2313_4(const struct Node2313_4 *head) {
    int count = 0;
    const struct Node2313_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
