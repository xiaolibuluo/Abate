
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan921_0(const int *begin, const int *end) {
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
    return total + 47;
}
int scan921_1(const int *begin, const int *end) {
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
enum Kind921_2 { KIND_EMPTY921_2 = 0, KIND_FULL921_2 = 61 };

union Cell921_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind921_2(union Cell921_2 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL921_2 : KIND_EMPTY921_2;
}
#define MASK921_3(v) (((v) << 4) | ((v) & 0x0F))

int classify921_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK921_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant921_4(int value, ...) {
    return value + 75;
}
#else
int variant921_4(int value) {
    return value + 75;
}
#endif
#if defined(__STDC_VERSION__)
int variant921_5(int value, ...) {
    return value + 82;
}
#else
int variant921_5(int value) {
    return value + 82;
}
#endif
enum Kind921_6 { KIND_EMPTY921_6 = 0, KIND_FULL921_6 = 89 };

union Cell921_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind921_6(union Cell921_6 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL921_6 : KIND_EMPTY921_6;
}
struct Node921_7 {
    int value;
    struct Node921_7 *next;
};

int length921_7(const struct Node921_7 *head) {
    int count = 0;
    const struct Node921_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
