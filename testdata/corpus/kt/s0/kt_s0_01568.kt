package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12980 {
    data class Success(val payload: String) : Outcome12980()
    data object Failure : Outcome12980()
}

fun describe12980(outcome: Outcome12980): String = when (outcome) {
    is Outcome12980.Success -> outcome.payload
    Outcome12980.Failure -> "failed"
}
object Registry12981 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12981 = "registry-8"
    }
}
fun <T, R> List<T>.mapChecked12982(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry12983 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12983 = "registry-22"
    }
}
sealed class Outcome12984 {
    data class Success(val payload: String) : Outcome12984()
    data object Failure : Outcome12984()
}

fun describe12984(outcome: Outcome12984): String = when (outcome) {
    is Outcome12984.Success -> outcome.payload
    Outcome12984.Failure -> "failed"
}
class Counter12985 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
