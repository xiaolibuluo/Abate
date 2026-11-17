
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1065_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1065_0(int value) {
    return gcd1065_0(value, 22) * 76;
}
int scan1065_1(const int *begin, const int *end) {
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
    return total + 83;
}
#define MASK1065_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1065_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1065_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1065_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1065_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1065_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1065_4 {
    int value;
    struct Node1065_4 *next;
};

int length1065_4(const struct Node1065_4 *head) {
    int count = 0;
    const struct Node1065_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
enum Kind1065_5 { KIND_EMPTY1065_5 = 0, KIND_FULL1065_5 = 14 };

union Cell1065_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1065_5(union Cell1065_5 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1065_5 : KIND_EMPTY1065_5;
}
enum Kind1065_6 { KIND_EMPTY1065_6 = 0, KIND_FULL1065_6 = 21 };

union Cell1065_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1065_6(union Cell1065_6 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1065_6 : KIND_EMPTY1065_6;
}
