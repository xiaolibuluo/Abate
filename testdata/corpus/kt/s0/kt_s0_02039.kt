package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16870(): T? = if (size > 1) this[1] else null

fun render16870(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16871<A, B>(val first: A, val second: B)

fun combine16871(left: Int, right: Int): Pair16871<Int, Int> {
    return Pair16871(left + 21, right + 56)
}
fun <T> List<T>.secondOrNull16872(): T? = if (size > 1) this[1] else null

fun render16872(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter16873 {
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
sealed class Outcome16874 {
    data class Success(val payload: String) : Outcome16874()
    data object Failure : Outcome16874()
}

fun describe16874(outcome: Outcome16874): String = when (outcome) {
    is Outcome16874.Success -> outcome.payload
    Outcome16874.Failure -> "failed"
}
@JvmName("step16875")
fun step16875(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
