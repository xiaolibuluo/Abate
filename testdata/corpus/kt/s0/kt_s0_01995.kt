package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16500")
fun step16500(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
fun <T> List<T>.secondOrNull16501(): T? = if (size > 1) this[1] else null

fun render16501(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry16502 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16502 = "registry-32"
    }
}
data class Pair16503<A, B>(val first: A, val second: B)

fun combine16503(left: Int, right: Int): Pair16503<Int, Int> {
    return Pair16503(left + 39, right + 63)
}
fun <T, R> List<T>.mapChecked16504(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome16505 {
    data class Success(val payload: String) : Outcome16505()
    data object Failure : Outcome16505()
}

fun describe16505(outcome: Outcome16505): String = when (outcome) {
    is Outcome16505.Success -> outcome.payload
    Outcome16505.Failure -> "failed"
}
class Counter16506 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
sealed class Outcome16507 {
    data class Success(val payload: String) : Outcome16507()
    data object Failure : Outcome16507()
}

fun describe16507(outcome: Outcome16507): String = when (outcome) {
    is Outcome16507.Success -> outcome.payload
    Outcome16507.Failure -> "failed"
}
