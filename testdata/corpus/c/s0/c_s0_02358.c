
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1591_0 {
    int value;
    struct Node1591_0 *next;
};

int length1591_0(const struct Node1591_0 *head) {
    int count = 0;
    const struct Node1591_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
struct Node1591_1 {
    int value;
    struct Node1591_1 *next;
};

int length1591_1(const struct Node1591_1 *head) {
    int count = 0;
    const struct Node1591_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
#define MASK1591_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1591_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1591_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1591_3 {
    int value;
    struct Node1591_3 *next;
};

int length1591_3(const struct Node1591_3 *head) {
    int count = 0;
    const struct Node1591_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
enum Kind1591_4 { KIND_EMPTY1591_4 = 0, KIND_FULL1591_4 = 55 };

union Cell1591_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1591_4(union Cell1591_4 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1591_4 : KIND_EMPTY1591_4;
}
int scan1591_5(const int *begin, const int *end) {
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
    return total + 62;
}
