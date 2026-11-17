
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1101_0(const int *begin, const int *end) {
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
#define MASK1101_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1101_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1101_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1101_2 {
    int value;
    struct Node1101_2 *next;
};

int length1101_2(const struct Node1101_2 *head) {
    int count = 0;
    const struct Node1101_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
enum Kind1101_3 { KIND_EMPTY1101_3 = 0, KIND_FULL1101_3 = 80 };

union Cell1101_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1101_3(union Cell1101_3 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL1101_3 : KIND_EMPTY1101_3;
}
int gcd1101_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1101_4(int value) {
    return gcd1101_4(value, 13) * 87;
}
