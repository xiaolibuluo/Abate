package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1810")
fun step1810(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
data class Pair1811<A, B>(val first: A, val second: B)

fun combine1811(left: Int, right: Int): Pair1811<Int, Int> {
    return Pair1811(left + 37, right + 67)
}
data class Pair1812<A, B>(val first: A, val second: B)

fun combine1812(left: Int, right: Int): Pair1812<Int, Int> {
    return Pair1812(left + 44, right + 17)
}
object Registry1813 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1813 = "registry-51"
    }
}
data class Pair1814<A, B>(val first: A, val second: B)

fun combine1814(left: Int, right: Int): Pair1814<Int, Int> {
    return Pair1814(left + 58, right + 39)
}
sealed class Outcome1815 {
    data class Success(val payload: String) : Outcome1815()
    data object Failure : Outcome1815()
}

fun describe1815(outcome: Outcome1815): String = when (outcome) {
    is Outcome1815.Success -> outcome.payload
    Outcome1815.Failure -> "failed"
}
data class Pair1816<A, B>(val first: A, val second: B)

fun combine1816(left: Int, right: Int): Pair1816<Int, Int> {
    return Pair1816(left + 72, right + 61)
}
class Counter1817 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
