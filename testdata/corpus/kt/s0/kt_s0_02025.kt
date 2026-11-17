package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16750(): T? = if (size > 1) this[1] else null

fun render16750(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16751 {
    data class Success(val payload: String) : Outcome16751()
    data object Failure : Outcome16751()
}

fun describe16751(outcome: Outcome16751): String = when (outcome) {
    is Outcome16751.Success -> outcome.payload
    Outcome16751.Failure -> "failed"
}
fun <T> List<T>.secondOrNull16752(): T? = if (size > 1) this[1] else null

fun render16752(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16753")
fun step16753(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
data class Pair16754<A, B>(val first: A, val second: B)

fun combine16754(left: Int, right: Int): Pair16754<Int, Int> {
    return Pair16754(left + 80, right + 66)
}
