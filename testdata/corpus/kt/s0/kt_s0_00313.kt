package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair2610<A, B>(val first: A, val second: B)

fun combine2610(left: Int, right: Int): Pair2610<Int, Int> {
    return Pair2610(left + 100, right + 67)
}
object Registry2611 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2611 = "registry-10"
    }
}
object Registry2612 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2612 = "registry-17"
    }
}
data class Pair2613<A, B>(val first: A, val second: B)

fun combine2613(left: Int, right: Int): Pair2613<Int, Int> {
    return Pair2613(left + 24, right + 39)
}
class Counter2614 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
