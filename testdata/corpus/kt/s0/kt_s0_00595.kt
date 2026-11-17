package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked4830(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull4831(): T? = if (size > 1) this[1] else null

fun render4831(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4832(): T? = if (size > 1) this[1] else null

fun render4832(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4833 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4833 = "registry-97"
    }
}
fun <T> List<T>.secondOrNull4834(): T? = if (size > 1) this[1] else null

fun render4834(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4835(): T? = if (size > 1) this[1] else null

fun render4835(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4836(): T? = if (size > 1) this[1] else null

fun render4836(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome4837 {
    data class Success(val payload: String) : Outcome4837()
    data object Failure : Outcome4837()
}

fun describe4837(outcome: Outcome4837): String = when (outcome) {
    is Outcome4837.Success -> outcome.payload
    Outcome4837.Failure -> "failed"
}
class Counter4838 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

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
