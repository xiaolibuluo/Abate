package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter16520 {
    var current: Int by Delegates.observable(44) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked16521(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16522<A, B>(val first: A, val second: B)

fun combine16522(left: Int, right: Int): Pair16522<Int, Int> {
    return Pair16522(left + 58, right + 66)
}
data class Pair16523<A, B>(val first: A, val second: B)

fun combine16523(left: Int, right: Int): Pair16523<Int, Int> {
    return Pair16523(left + 65, right + 16)
}
object Registry16524 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16524 = "registry-72"
    }
}
sealed class Outcome16525 {
    data class Success(val payload: String) : Outcome16525()
    data object Failure : Outcome16525()
}

fun describe16525(outcome: Outcome16525): String = when (outcome) {
    is Outcome16525.Success -> outcome.payload
    Outcome16525.Failure -> "failed"
}
sealed class Outcome16526 {
    data class Success(val payload: String) : Outcome16526()
    data object Failure : Outcome16526()
}

fun describe16526(outcome: Outcome16526): String = when (outcome) {
    is Outcome16526.Success -> outcome.payload
    Outcome16526.Failure -> "failed"
}
fun <T> List<T>.secondOrNull16527(): T? = if (size > 1) this[1] else null

fun render16527(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
