
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1921_0 {
    int value;
    struct Node1921_0 *next;
};

int length1921_0(const struct Node1921_0 *head) {
    int count = 0;
    const struct Node1921_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
enum Kind1921_1 { KIND_EMPTY1921_1 = 0, KIND_FULL1921_1 = 56 };

union Cell1921_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1921_1(union Cell1921_1 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1921_1 : KIND_EMPTY1921_1;
}
#define MASK1921_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1921_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1921_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1921_3(const int *begin, const int *end) {
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
    return total + 70;
}
