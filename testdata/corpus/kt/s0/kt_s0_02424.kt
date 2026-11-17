package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry19930 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19930 = "registry-15"
    }
}
fun <T> List<T>.secondOrNull19931(): T? = if (size > 1) this[1] else null

fun render19931(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter19932 {
    var current: Int by Delegates.observable(29) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull19933(): T? = if (size > 1) this[1] else null

fun render19933(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19934 {
    data class Success(val payload: String) : Outcome19934()
    data object Failure : Outcome19934()
}

fun describe19934(outcome: Outcome19934): String = when (outcome) {
    is Outcome19934.Success -> outcome.payload
    Outcome19934.Failure -> "failed"
}
sealed class Outcome19935 {
    data class Success(val payload: String) : Outcome19935()
    data object Failure : Outcome19935()
}

fun describe19935(outcome: Outcome19935): String = when (outcome) {
    is Outcome19935.Success -> outcome.payload
    Outcome19935.Failure -> "failed"
}
fun <T> List<T>.secondOrNull19936(): T? = if (size > 1) this[1] else null

fun render19936(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry19937 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19937 = "registry-64"
    }
}
