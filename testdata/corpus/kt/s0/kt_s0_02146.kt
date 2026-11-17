package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17710 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17710 = "registry-39"
    }
}
data class Pair17711<A, B>(val first: A, val second: B)

fun combine17711(left: Int, right: Int): Pair17711<Int, Int> {
    return Pair17711(left + 46, right + 34)
}
object Registry17712 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17712 = "registry-53"
    }
}
class Counter17713 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
