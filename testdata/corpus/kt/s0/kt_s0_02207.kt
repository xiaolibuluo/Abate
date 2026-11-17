package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18180")
fun step18180(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 68
}
object Registry18181 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18181 = "registry-75"
    }
}
class Counter18182 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
data class Pair18183<A, B>(val first: A, val second: B)

fun combine18183(left: Int, right: Int): Pair18183<Int, Int> {
    return Pair18183(left + 89, right + 19)
}
