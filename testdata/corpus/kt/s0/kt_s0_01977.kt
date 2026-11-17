package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked16340(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter16341 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
@JvmName("step16342")
fun step16342(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
data class Pair16343<A, B>(val first: A, val second: B)

fun combine16343(left: Int, right: Int): Pair16343<Int, Int> {
    return Pair16343(left + 25, right + 12)
}
data class Pair16344<A, B>(val first: A, val second: B)

fun combine16344(left: Int, right: Int): Pair16344<Int, Int> {
    return Pair16344(left + 32, right + 23)
}
fun <T, R> List<T>.mapChecked16345(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
