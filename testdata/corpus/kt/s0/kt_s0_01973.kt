package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry16310 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16310 = "registry-62"
    }
}
fun <T> List<T>.secondOrNull16311(): T? = if (size > 1) this[1] else null

fun render16311(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16312<A, B>(val first: A, val second: B)

fun combine16312(left: Int, right: Int): Pair16312<Int, Int> {
    return Pair16312(left + 76, right + 41)
}
data class Pair16313<A, B>(val first: A, val second: B)

fun combine16313(left: Int, right: Int): Pair16313<Int, Int> {
    return Pair16313(left + 83, right + 52)
}
@JvmName("step16314")
fun step16314(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
fun <T> List<T>.secondOrNull16315(): T? = if (size > 1) this[1] else null

fun render16315(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
