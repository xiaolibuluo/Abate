package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8650 {
    data class Success(val payload: String) : Outcome8650()
    data object Failure : Outcome8650()
}

fun describe8650(outcome: Outcome8650): String = when (outcome) {
    is Outcome8650.Success -> outcome.payload
    Outcome8650.Failure -> "failed"
}
sealed class Outcome8651 {
    data class Success(val payload: String) : Outcome8651()
    data object Failure : Outcome8651()
}

fun describe8651(outcome: Outcome8651): String = when (outcome) {
    is Outcome8651.Success -> outcome.payload
    Outcome8651.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked8652(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter8653 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
object Registry8654 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8654 = "registry-26"
    }
}
sealed class Outcome8655 {
    data class Success(val payload: String) : Outcome8655()
    data object Failure : Outcome8655()
}

fun describe8655(outcome: Outcome8655): String = when (outcome) {
    is Outcome8655.Success -> outcome.payload
    Outcome8655.Failure -> "failed"
}
