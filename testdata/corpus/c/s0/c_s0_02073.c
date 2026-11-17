
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1397_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1397_0(int value) {
    return gcd1397_0(value, 28) * 27;
}
enum Kind1397_1 { KIND_EMPTY1397_1 = 0, KIND_FULL1397_1 = 34 };

union Cell1397_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1397_1(union Cell1397_1 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1397_1 : KIND_EMPTY1397_1;
}
#define MASK1397_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1397_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1397_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1397_3(const int *begin, const int *end) {
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
    return total + 48;
}
enum Kind1397_4 { KIND_EMPTY1397_4 = 0, KIND_FULL1397_4 = 55 };

union Cell1397_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1397_4(union Cell1397_4 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1397_4 : KIND_EMPTY1397_4;
}
struct Node1397_5 {
    int value;
    struct Node1397_5 *next;
};

int length1397_5(const struct Node1397_5 *head) {
    int count = 0;
    const struct Node1397_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
