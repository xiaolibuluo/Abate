package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1280")
fun step1280(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
object Registry1281 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1281 = "registry-27"
    }
}
fun <T, R> List<T>.mapChecked1282(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair1283<A, B>(val first: A, val second: B)

fun combine1283(left: Int, right: Int): Pair1283<Int, Int> {
    return Pair1283(left + 41, right + 23)
}
class Counter1284 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked1285(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair1286<A, B>(val first: A, val second: B)

fun combine1286(left: Int, right: Int): Pair1286<Int, Int> {
    return Pair1286(left + 62, right + 56)
}
