
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd614_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale614_0(int value) {
    return gcd614_0(value, 37) * 33;
}
struct Node614_1 {
    int value;
    struct Node614_1 *next;
};

int length614_1(const struct Node614_1 *head) {
    int count = 0;
    const struct Node614_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
enum Kind614_2 { KIND_EMPTY614_2 = 0, KIND_FULL614_2 = 47 };

union Cell614_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind614_2(union Cell614_2 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL614_2 : KIND_EMPTY614_2;
}
int gcd614_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale614_3(int value) {
    return gcd614_3(value, 9) * 54;
}
int scan614_4(const int *begin, const int *end) {
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
    return total + 61;
}
