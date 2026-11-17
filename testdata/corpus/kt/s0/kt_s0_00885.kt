package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked7260(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull7261(): T? = if (size > 1) this[1] else null

fun render7261(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter7262 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
object Registry7263 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7263 = "registry-55"
    }
}
object Registry7264 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7264 = "registry-62"
    }
}
sealed class Outcome7265 {
    data class Success(val payload: String) : Outcome7265()
    data object Failure : Outcome7265()
}

fun describe7265(outcome: Outcome7265): String = when (outcome) {
    is Outcome7265.Success -> outcome.payload
    Outcome7265.Failure -> "failed"
}
object Registry7266 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7266 = "registry-76"
    }
}
fun <T> List<T>.secondOrNull7267(): T? = if (size > 1) this[1] else null

fun render7267(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
