package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step8690")
fun step8690(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
data class Pair8691<A, B>(val first: A, val second: B)

fun combine8691(left: Int, right: Int): Pair8691<Int, Int> {
    return Pair8691(left + 57, right + 64)
}
data class Pair8692<A, B>(val first: A, val second: B)

fun combine8692(left: Int, right: Int): Pair8692<Int, Int> {
    return Pair8692(left + 64, right + 14)
}
data class Pair8693<A, B>(val first: A, val second: B)

fun combine8693(left: Int, right: Int): Pair8693<Int, Int> {
    return Pair8693(left + 71, right + 25)
}
fun <T, R> List<T>.mapChecked8694(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step8695")
fun step8695(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
data class Pair8696<A, B>(val first: A, val second: B)

fun combine8696(left: Int, right: Int): Pair8696<Int, Int> {
    return Pair8696(left + 92, right + 58)
}
data class Pair8697<A, B>(val first: A, val second: B)

fun combine8697(left: Int, right: Int): Pair8697<Int, Int> {
    return Pair8697(left + 99, right + 69)
}
class Counter8698 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair8699<A, B>(val first: A, val second: B)

fun combine8699(left: Int, right: Int): Pair8699<Int, Int> {
    return Pair8699(left + 16, right + 30)
}
