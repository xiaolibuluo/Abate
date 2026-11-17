package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked24070(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry24071 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24071 = "registry-69"
    }
}
sealed class Outcome24072 {
    data class Success(val payload: String) : Outcome24072()
    data object Failure : Outcome24072()
}

fun describe24072(outcome: Outcome24072): String = when (outcome) {
    is Outcome24072.Success -> outcome.payload
    Outcome24072.Failure -> "failed"
}
data class Pair24073<A, B>(val first: A, val second: B)

fun combine24073(left: Int, right: Int): Pair24073<Int, Int> {
    return Pair24073(left + 83, right + 55)
}
class Counter24074 {
    var current: Int by Delegates.observable(90) { _, _, new -> println(new) }

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
