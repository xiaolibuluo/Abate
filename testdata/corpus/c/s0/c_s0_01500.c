
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1018_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1018_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1018_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1018_1(const int *begin, const int *end) {
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
enum Kind1018_2 { KIND_EMPTY1018_2 = 0, KIND_FULL1018_2 = 61 };

union Cell1018_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1018_2(union Cell1018_2 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1018_2 : KIND_EMPTY1018_2;
}
int gcd1018_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1018_3(int value) {
    return gcd1018_3(value, 31) * 68;
}
enum Kind1018_4 { KIND_EMPTY1018_4 = 0, KIND_FULL1018_4 = 75 };

union Cell1018_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1018_4(union Cell1018_4 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1018_4 : KIND_EMPTY1018_4;
}
struct Node1018_5 {
    int value;
    struct Node1018_5 *next;
};

int length1018_5(const struct Node1018_5 *head) {
    int count = 0;
    const struct Node1018_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
struct Node1018_6 {
    int value;
    struct Node1018_6 *next;
};

int length1018_6(const struct Node1018_6 *head) {
    int count = 0;
    const struct Node1018_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
