package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7740 {
    data class Success(val payload: String) : Outcome7740()
    data object Failure : Outcome7740()
}

fun describe7740(outcome: Outcome7740): String = when (outcome) {
    is Outcome7740.Success -> outcome.payload
    Outcome7740.Failure -> "failed"
}
class Counter7741 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

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
object Registry7742 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7742 = "registry-90"
    }
}
fun <T, R> List<T>.mapChecked7743(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull7744(): T? = if (size > 1) this[1] else null

fun render7744(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter7745 {
    var current: Int by Delegates.observable(14) { _, _, new -> println(new) }

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
object Registry7746 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7746 = "registry-21"
    }
}
sealed class Outcome7747 {
    data class Success(val payload: String) : Outcome7747()
    data object Failure : Outcome7747()
}

fun describe7747(outcome: Outcome7747): String = when (outcome) {
    is Outcome7747.Success -> outcome.payload
    Outcome7747.Failure -> "failed"
}
