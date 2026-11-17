package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19260<A, B>(val first: A, val second: B)

fun combine19260(left: Int, right: Int): Pair19260<Int, Int> {
    return Pair19260(left + 17, right + 10)
}
fun <T, R> List<T>.mapChecked19261(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19262<A, B>(val first: A, val second: B)

fun combine19262(left: Int, right: Int): Pair19262<Int, Int> {
    return Pair19262(left + 31, right + 32)
}
class Counter19263 {
    var current: Int by Delegates.observable(38) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked19264(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step19265")
fun step19265(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
