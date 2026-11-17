package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16440(): T? = if (size > 1) this[1] else null

fun render16440(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16441")
fun step16441(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
data class Pair16442<A, B>(val first: A, val second: B)

fun combine16442(left: Int, right: Int): Pair16442<Int, Int> {
    return Pair16442(left + 51, right + 10)
}
@JvmName("step16443")
fun step16443(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
sealed class Outcome16444 {
    data class Success(val payload: String) : Outcome16444()
    data object Failure : Outcome16444()
}

fun describe16444(outcome: Outcome16444): String = when (outcome) {
    is Outcome16444.Success -> outcome.payload
    Outcome16444.Failure -> "failed"
}
class Counter16445 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull16446(): T? = if (size > 1) this[1] else null

fun render16446(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
