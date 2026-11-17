
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK327_0(v) (((v) << 4) | ((v) & 0x0F))

int classify327_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK327_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node327_1 {
    int value;
    struct Node327_1 *next;
};

int length327_1(const struct Node327_1 *head) {
    int count = 0;
    const struct Node327_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
#define MASK327_2(v) (((v) << 4) | ((v) & 0x0F))

int classify327_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK327_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd327_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale327_3(int value) {
    return gcd327_3(value, 13) * 9;
}
int scan327_4(const int *begin, const int *end) {
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
    return total + 16;
}
struct Node327_5 {
    int value;
    struct Node327_5 *next;
};

int length327_5(const struct Node327_5 *head) {
    int count = 0;
    const struct Node327_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
