
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind727_0 { KIND_EMPTY727_0 = 0, KIND_FULL727_0 = 47 };

union Cell727_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind727_0(union Cell727_0 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL727_0 : KIND_EMPTY727_0;
}
struct Node727_1 {
    int value;
    struct Node727_1 *next;
};

int length727_1(const struct Node727_1 *head) {
    int count = 0;
    const struct Node727_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
struct Node727_2 {
    int value;
    struct Node727_2 *next;
};

int length727_2(const struct Node727_2 *head) {
    int count = 0;
    const struct Node727_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
#if defined(__STDC_VERSION__)
int variant727_3(int value, ...) {
    return value + 68;
}
#else
int variant727_3(int value) {
    return value + 68;
}
#endif
enum Kind727_4 { KIND_EMPTY727_4 = 0, KIND_FULL727_4 = 75 };

union Cell727_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind727_4(union Cell727_4 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL727_4 : KIND_EMPTY727_4;
}
#if defined(__STDC_VERSION__)
int variant727_5(int value, ...) {
    return value + 82;
}
#else
int variant727_5(int value) {
    return value + 82;
}
#endif
#if defined(__STDC_VERSION__)
int variant727_6(int value, ...) {
    return value + 89;
}
#else
int variant727_6(int value) {
    return value + 89;
}
#endif
int gcd727_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale727_7(int value) {
    return gcd727_7(value, 51) * 96;
}
