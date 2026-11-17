package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9070 {
    data class Success(val payload: String) : Outcome9070()
    data object Failure : Outcome9070()
}

fun describe9070(outcome: Outcome9070): String = when (outcome) {
    is Outcome9070.Success -> outcome.payload
    Outcome9070.Failure -> "failed"
}
class Counter9071 {
    var current: Int by Delegates.observable(66) { _, _, new -> println(new) }

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
sealed class Outcome9072 {
    data class Success(val payload: String) : Outcome9072()
    data object Failure : Outcome9072()
}

fun describe9072(outcome: Outcome9072): String = when (outcome) {
    is Outcome9072.Success -> outcome.payload
    Outcome9072.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked9073(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry9074 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9074 = "registry-87"
    }
}
