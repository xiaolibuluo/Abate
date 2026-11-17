package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18300 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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
object Registry18301 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18301 = "registry-37"
    }
}
data class Pair18302<A, B>(val first: A, val second: B)

fun combine18302(left: Int, right: Int): Pair18302<Int, Int> {
    return Pair18302(left + 44, right + 31)
}
data class Pair18303<A, B>(val first: A, val second: B)

fun combine18303(left: Int, right: Int): Pair18303<Int, Int> {
    return Pair18303(left + 51, right + 42)
}
fun <T, R> List<T>.mapChecked18304(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair18305<A, B>(val first: A, val second: B)

fun combine18305(left: Int, right: Int): Pair18305<Int, Int> {
    return Pair18305(left + 65, right + 64)
}
