package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19770<A, B>(val first: A, val second: B)

fun combine19770(left: Int, right: Int): Pair19770<Int, Int> {
    return Pair19770(left + 98, right + 62)
}
object Registry19771 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19771 = "registry-8"
    }
}
data class Pair19772<A, B>(val first: A, val second: B)

fun combine19772(left: Int, right: Int): Pair19772<Int, Int> {
    return Pair19772(left + 15, right + 23)
}
class Counter19773 {
    var current: Int by Delegates.observable(22) { _, _, new -> println(new) }

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
sealed class Outcome19774 {
    data class Success(val payload: String) : Outcome19774()
    data object Failure : Outcome19774()
}

fun describe19774(outcome: Outcome19774): String = when (outcome) {
    is Outcome19774.Success -> outcome.payload
    Outcome19774.Failure -> "failed"
}
sealed class Outcome19775 {
    data class Success(val payload: String) : Outcome19775()
    data object Failure : Outcome19775()
}

fun describe19775(outcome: Outcome19775): String = when (outcome) {
    is Outcome19775.Success -> outcome.payload
    Outcome19775.Failure -> "failed"
}
