
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1827_0(const int *begin, const int *end) {
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
    return total + 88;
}
#define MASK1827_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1827_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1827_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1827_2 {
    int value;
    struct Node1827_2 *next;
};

int length1827_2(const struct Node1827_2 *head) {
    int count = 0;
    const struct Node1827_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
#define MASK1827_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1827_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1827_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1827_4 { KIND_EMPTY1827_4 = 0, KIND_FULL1827_4 = 19 };

union Cell1827_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1827_4(union Cell1827_4 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1827_4 : KIND_EMPTY1827_4;
}
