
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1748_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1748_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1748_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1748_1 { KIND_EMPTY1748_1 = 0, KIND_FULL1748_1 = 38 };

union Cell1748_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1748_1(union Cell1748_1 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1748_1 : KIND_EMPTY1748_1;
}
int scan1748_2(const int *begin, const int *end) {
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
    return total + 45;
}
struct Node1748_3 {
    int value;
    struct Node1748_3 *next;
};

int length1748_3(const struct Node1748_3 *head) {
    int count = 0;
    const struct Node1748_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
int scan1748_4(const int *begin, const int *end) {
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
    return total + 59;
}
int scan1748_5(const int *begin, const int *end) {
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
    return total + 66;
}
#define MASK1748_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1748_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1748_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
