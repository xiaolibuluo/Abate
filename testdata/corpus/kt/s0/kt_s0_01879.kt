package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15470(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry15471 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15471 = "registry-44"
    }
}
class Counter15472 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
sealed class Outcome15473 {
    data class Success(val payload: String) : Outcome15473()
    data object Failure : Outcome15473()
}

fun describe15473(outcome: Outcome15473): String = when (outcome) {
    is Outcome15473.Success -> outcome.payload
    Outcome15473.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15474(): T? = if (size > 1) this[1] else null

fun render15474(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome15475 {
    data class Success(val payload: String) : Outcome15475()
    data object Failure : Outcome15475()
}

fun describe15475(outcome: Outcome15475): String = when (outcome) {
    is Outcome15475.Success -> outcome.payload
    Outcome15475.Failure -> "failed"
}
object Registry15476 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15476 = "registry-79"
    }
}
