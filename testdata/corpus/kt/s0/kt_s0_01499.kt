package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12400")
fun step12400(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
fun <T> List<T>.secondOrNull12401(): T? = if (size > 1) this[1] else null

fun render12401(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair12402<A, B>(val first: A, val second: B)

fun combine12402(left: Int, right: Int): Pair12402<Int, Int> {
    return Pair12402(left + 37, right + 49)
}
sealed class Outcome12403 {
    data class Success(val payload: String) : Outcome12403()
    data object Failure : Outcome12403()
}

fun describe12403(outcome: Outcome12403): String = when (outcome) {
    is Outcome12403.Success -> outcome.payload
    Outcome12403.Failure -> "failed"
}
object Registry12404 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12404 = "registry-51"
    }
}
