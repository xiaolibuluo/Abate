package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23680 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23680 = "registry-40"
    }
}
object Registry23681 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23681 = "registry-47"
    }
}
data class Pair23682<A, B>(val first: A, val second: B)

fun combine23682(left: Int, right: Int): Pair23682<Int, Int> {
    return Pair23682(left + 54, right + 15)
}
sealed class Outcome23683 {
    data class Success(val payload: String) : Outcome23683()
    data object Failure : Outcome23683()
}

fun describe23683(outcome: Outcome23683): String = when (outcome) {
    is Outcome23683.Success -> outcome.payload
    Outcome23683.Failure -> "failed"
}
class Counter23684 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
sealed class Outcome23685 {
    data class Success(val payload: String) : Outcome23685()
    data object Failure : Outcome23685()
}

fun describe23685(outcome: Outcome23685): String = when (outcome) {
    is Outcome23685.Success -> outcome.payload
    Outcome23685.Failure -> "failed"
}
