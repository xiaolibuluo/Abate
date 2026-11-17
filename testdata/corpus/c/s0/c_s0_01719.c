
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1155_0 {
    int value;
    struct Node1155_0 *next;
};

int length1155_0(const struct Node1155_0 *head) {
    int count = 0;
    const struct Node1155_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
struct Node1155_1 {
    int value;
    struct Node1155_1 *next;
};

int length1155_1(const struct Node1155_1 *head) {
    int count = 0;
    const struct Node1155_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
#define MASK1155_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1155_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1155_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1155_3 { KIND_EMPTY1155_3 = 0, KIND_FULL1155_3 = 6 };

union Cell1155_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1155_3(union Cell1155_3 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1155_3 : KIND_EMPTY1155_3;
}
int scan1155_4(const int *begin, const int *end) {
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
    return total + 13;
}
#if defined(__STDC_VERSION__)
int variant1155_5(int value, ...) {
    return value + 20;
}
#else
int variant1155_5(int value) {
    return value + 20;
}
#endif
