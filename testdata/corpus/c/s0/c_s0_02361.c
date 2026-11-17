
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1592_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1592_0(int value) {
    return gcd1592_0(value, 51) * 40;
}
int scan1592_1(const int *begin, const int *end) {
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
struct Node1592_2 {
    int value;
    struct Node1592_2 *next;
};

int length1592_2(const struct Node1592_2 *head) {
    int count = 0;
    const struct Node1592_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
enum Kind1592_3 { KIND_EMPTY1592_3 = 0, KIND_FULL1592_3 = 61 };

union Cell1592_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1592_3(union Cell1592_3 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1592_3 : KIND_EMPTY1592_3;
}
#define MASK1592_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1592_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1592_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1592_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1592_5(int value) {
    return gcd1592_5(value, 45) * 75;
}
