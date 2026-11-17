package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter13090 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
object Registry13091 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13091 = "registry-54"
    }
}
fun <T> List<T>.secondOrNull13092(): T? = if (size > 1) this[1] else null

fun render13092(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13093 {
    data class Success(val payload: String) : Outcome13093()
    data object Failure : Outcome13093()
}

fun describe13093(outcome: Outcome13093): String = when (outcome) {
    is Outcome13093.Success -> outcome.payload
    Outcome13093.Failure -> "failed"
}
sealed class Outcome13094 {
    data class Success(val payload: String) : Outcome13094()
    data object Failure : Outcome13094()
}

fun describe13094(outcome: Outcome13094): String = when (outcome) {
    is Outcome13094.Success -> outcome.payload
    Outcome13094.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13095(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome13096 {
    data class Success(val payload: String) : Outcome13096()
    data object Failure : Outcome13096()
}

fun describe13096(outcome: Outcome13096): String = when (outcome) {
    is Outcome13096.Success -> outcome.payload
    Outcome13096.Failure -> "failed"
}
