package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry22630 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22630 = "registry-33"
    }
}
sealed class Outcome22631 {
    data class Success(val payload: String) : Outcome22631()
    data object Failure : Outcome22631()
}

fun describe22631(outcome: Outcome22631): String = when (outcome) {
    is Outcome22631.Success -> outcome.payload
    Outcome22631.Failure -> "failed"
}
@JvmName("step22632")
fun step22632(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
data class Pair22633<A, B>(val first: A, val second: B)

fun combine22633(left: Int, right: Int): Pair22633<Int, Int> {
    return Pair22633(left + 54, right + 23)
}
fun <T> List<T>.secondOrNull22634(): T? = if (size > 1) this[1] else null

fun render22634(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22635 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22635 = "registry-68"
    }
}
