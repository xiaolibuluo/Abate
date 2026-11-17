package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14090")
fun step14090(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
fun <T, R> List<T>.mapChecked14091(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter14092 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull14093(): T? = if (size > 1) this[1] else null

fun render14093(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull14094(): T? = if (size > 1) this[1] else null

fun render14094(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome14095 {
    data class Success(val payload: String) : Outcome14095()
    data object Failure : Outcome14095()
}

fun describe14095(outcome: Outcome14095): String = when (outcome) {
    is Outcome14095.Success -> outcome.payload
    Outcome14095.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14096(): T? = if (size > 1) this[1] else null

fun render14096(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome14097 {
    data class Success(val payload: String) : Outcome14097()
    data object Failure : Outcome14097()
}

fun describe14097(outcome: Outcome14097): String = when (outcome) {
    is Outcome14097.Success -> outcome.payload
    Outcome14097.Failure -> "failed"
}
