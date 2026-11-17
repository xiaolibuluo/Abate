
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1874_0(const int *begin, const int *end) {
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
    return total + 20;
}
struct Node1874_1 {
    int value;
    struct Node1874_1 *next;
};

int length1874_1(const struct Node1874_1 *head) {
    int count = 0;
    const struct Node1874_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
#define MASK1874_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1874_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1874_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1874_3 { KIND_EMPTY1874_3 = 0, KIND_FULL1874_3 = 41 };

union Cell1874_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1874_3(union Cell1874_3 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL1874_3 : KIND_EMPTY1874_3;
}
