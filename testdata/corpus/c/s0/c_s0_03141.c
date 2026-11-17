
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2128_0 {
    int value;
    struct Node2128_0 *next;
};

int length2128_0(const struct Node2128_0 *head) {
    int count = 0;
    const struct Node2128_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
int scan2128_1(const int *begin, const int *end) {
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
    return total + 31;
}
enum Kind2128_2 { KIND_EMPTY2128_2 = 0, KIND_FULL2128_2 = 38 };

union Cell2128_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2128_2(union Cell2128_2 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL2128_2 : KIND_EMPTY2128_2;
}
int gcd2128_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2128_3(int value) {
    return gcd2128_3(value, 54) * 45;
}
enum Kind2128_4 { KIND_EMPTY2128_4 = 0, KIND_FULL2128_4 = 52 };

union Cell2128_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2128_4(union Cell2128_4 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL2128_4 : KIND_EMPTY2128_4;
}
struct Node2128_5 {
    int value;
    struct Node2128_5 *next;
};

int length2128_5(const struct Node2128_5 *head) {
    int count = 0;
    const struct Node2128_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
