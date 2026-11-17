package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry2660 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2660 = "registry-68"
    }
}
data class Pair2661<A, B>(val first: A, val second: B)

fun combine2661(left: Int, right: Int): Pair2661<Int, Int> {
    return Pair2661(left + 75, right + 52)
}
data class Pair2662<A, B>(val first: A, val second: B)

fun combine2662(left: Int, right: Int): Pair2662<Int, Int> {
    return Pair2662(left + 82, right + 63)
}
class Counter2663 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
@JvmName("step2664")
fun step2664(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
