
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1542_0 {
    int value;
    struct Node1542_0 *next;
};

int length1542_0(const struct Node1542_0 *head) {
    int count = 0;
    const struct Node1542_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
int scan1542_1(const int *begin, const int *end) {
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
    return total + 76;
}
struct Node1542_2 {
    int value;
    struct Node1542_2 *next;
};

int length1542_2(const struct Node1542_2 *head) {
    int count = 0;
    const struct Node1542_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
enum Kind1542_3 { KIND_EMPTY1542_3 = 0, KIND_FULL1542_3 = 90 };

union Cell1542_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1542_3(union Cell1542_3 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1542_3 : KIND_EMPTY1542_3;
}
int gcd1542_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1542_4(int value) {
    return gcd1542_4(value, 50) * 97;
}
