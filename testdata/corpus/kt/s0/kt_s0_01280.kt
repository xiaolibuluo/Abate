package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull10580(): T? = if (size > 1) this[1] else null

fun render10580(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome10581 {
    data class Success(val payload: String) : Outcome10581()
    data object Failure : Outcome10581()
}

fun describe10581(outcome: Outcome10581): String = when (outcome) {
    is Outcome10581.Success -> outcome.payload
    Outcome10581.Failure -> "failed"
}
object Registry10582 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10582 = "registry-96"
    }
}
fun <T, R> List<T>.mapChecked10583(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter10584 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

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
class Counter10585 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
