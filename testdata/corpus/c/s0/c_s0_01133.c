
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node777_0 {
    int value;
    struct Node777_0 *next;
};

int length777_0(const struct Node777_0 *head) {
    int count = 0;
    const struct Node777_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
enum Kind777_1 { KIND_EMPTY777_1 = 0, KIND_FULL777_1 = 25 };

union Cell777_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind777_1(union Cell777_1 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL777_1 : KIND_EMPTY777_1;
}
#define MASK777_2(v) (((v) << 4) | ((v) & 0x0F))

int classify777_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK777_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node777_3 {
    int value;
    struct Node777_3 *next;
};

int length777_3(const struct Node777_3 *head) {
    int count = 0;
    const struct Node777_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
#if defined(__STDC_VERSION__)
int variant777_4(int value, ...) {
    return value + 46;
}
#else
int variant777_4(int value) {
    return value + 46;
}
#endif
#if defined(__STDC_VERSION__)
int variant777_5(int value, ...) {
    return value + 53;
}
#else
int variant777_5(int value) {
    return value + 53;
}
#endif
#if defined(__STDC_VERSION__)
int variant777_6(int value, ...) {
    return value + 60;
}
#else
int variant777_6(int value) {
    return value + 60;
}
#endif
struct Node777_7 {
    int value;
    struct Node777_7 *next;
};

int length777_7(const struct Node777_7 *head) {
    int count = 0;
    const struct Node777_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
int gcd777_8(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale777_8(int value) {
    return gcd777_8(value, 46) * 74;
}
