package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10780 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10780 = "registry-51"
    }
}
data class Pair10781<A, B>(val first: A, val second: B)

fun combine10781(left: Int, right: Int): Pair10781<Int, Int> {
    return Pair10781(left + 58, right + 63)
}
class Counter10782 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
