package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull13320(): T? = if (size > 1) this[1] else null

fun render13320(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair13321<A, B>(val first: A, val second: B)

fun combine13321(left: Int, right: Int): Pair13321<Int, Int> {
    return Pair13321(left + 62, right + 11)
}
object Registry13322 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13322 = "registry-69"
    }
}
sealed class Outcome13323 {
    data class Success(val payload: String) : Outcome13323()
    data object Failure : Outcome13323()
}

fun describe13323(outcome: Outcome13323): String = when (outcome) {
    is Outcome13323.Success -> outcome.payload
    Outcome13323.Failure -> "failed"
}
