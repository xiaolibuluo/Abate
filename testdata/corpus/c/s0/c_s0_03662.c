
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2473_0 {
    int value;
    struct Node2473_0 *next;
};

int length2473_0(const struct Node2473_0 *head) {
    int count = 0;
    const struct Node2473_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
struct Node2473_1 {
    int value;
    struct Node2473_1 *next;
};

int length2473_1(const struct Node2473_1 *head) {
    int count = 0;
    const struct Node2473_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
int gcd2473_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2473_2(int value) {
    return gcd2473_2(value, 18) * 61;
}
int scan2473_3(const int *begin, const int *end) {
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
    return total + 68;
}
int scan2473_4(const int *begin, const int *end) {
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
    return total + 75;
}
enum Kind2473_5 { KIND_EMPTY2473_5 = 0, KIND_FULL2473_5 = 82 };

union Cell2473_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2473_5(union Cell2473_5 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL2473_5 : KIND_EMPTY2473_5;
}
