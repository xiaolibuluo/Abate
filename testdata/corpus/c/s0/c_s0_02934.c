
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1983_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1983_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1983_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1983_1(const int *begin, const int *end) {
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
    return total + 86;
}
int gcd1983_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1983_2(int value) {
    return gcd1983_2(value, 65) * 93;
}
struct Node1983_3 {
    int value;
    struct Node1983_3 *next;
};

int length1983_3(const struct Node1983_3 *head) {
    int count = 0;
    const struct Node1983_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
struct Node1983_4 {
    int value;
    struct Node1983_4 *next;
};

int length1983_4(const struct Node1983_4 *head) {
    int count = 0;
    const struct Node1983_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
enum Kind1983_5 { KIND_EMPTY1983_5 = 0, KIND_FULL1983_5 = 17 };

union Cell1983_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1983_5(union Cell1983_5 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1983_5 : KIND_EMPTY1983_5;
}
