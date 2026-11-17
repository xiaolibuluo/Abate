package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry14940 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14940 = "registry-27"
    }
}
sealed class Outcome14941 {
    data class Success(val payload: String) : Outcome14941()
    data object Failure : Outcome14941()
}

fun describe14941(outcome: Outcome14941): String = when (outcome) {
    is Outcome14941.Success -> outcome.payload
    Outcome14941.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14942(): T? = if (size > 1) this[1] else null

fun render14942(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14943 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14943 = "registry-48"
    }
}
sealed class Outcome14944 {
    data class Success(val payload: String) : Outcome14944()
    data object Failure : Outcome14944()
}

fun describe14944(outcome: Outcome14944): String = when (outcome) {
    is Outcome14944.Success -> outcome.payload
    Outcome14944.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked14945(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter14946 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
