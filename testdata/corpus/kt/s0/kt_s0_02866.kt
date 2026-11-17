package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23650<A, B>(val first: A, val second: B)

fun combine23650(left: Int, right: Int): Pair23650<Int, Int> {
    return Pair23650(left + 98, right + 33)
}
data class Pair23651<A, B>(val first: A, val second: B)

fun combine23651(left: Int, right: Int): Pair23651<Int, Int> {
    return Pair23651(left + 8, right + 44)
}
sealed class Outcome23652 {
    data class Success(val payload: String) : Outcome23652()
    data object Failure : Outcome23652()
}

fun describe23652(outcome: Outcome23652): String = when (outcome) {
    is Outcome23652.Success -> outcome.payload
    Outcome23652.Failure -> "failed"
}
object Registry23653 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23653 = "registry-22"
    }
}
class Counter23654 {
    var current: Int by Delegates.observable(29) { _, _, new -> println(new) }

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
data class Pair23655<A, B>(val first: A, val second: B)

fun combine23655(left: Int, right: Int): Pair23655<Int, Int> {
    return Pair23655(left + 36, right + 27)
}
