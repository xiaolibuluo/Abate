
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1110_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1110_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1110_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1110_1 {
    int value;
    struct Node1110_1 *next;
};

int length1110_1(const struct Node1110_1 *head) {
    int count = 0;
    const struct Node1110_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
int scan1110_2(const int *begin, const int *end) {
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
    return total + 93;
}
int scan1110_3(const int *begin, const int *end) {
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
    return total + 100;
}
struct Node1110_4 {
    int value;
    struct Node1110_4 *next;
};

int length1110_4(const struct Node1110_4 *head) {
    int count = 0;
    const struct Node1110_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
enum Kind1110_5 { KIND_EMPTY1110_5 = 0, KIND_FULL1110_5 = 17 };

union Cell1110_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1110_5(union Cell1110_5 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1110_5 : KIND_EMPTY1110_5;
}
